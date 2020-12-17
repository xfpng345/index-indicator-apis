import React, { useEffect, useState } from 'react'
import { CntWrap } from '../UiKits/index';

const Fgi = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        async function fetchData() {
            const resResult = await fetch('/server/api/fgi/')
            resResult
                .json()
                .then((resResult) => setData(resResult))
                .catch(() => null)
        }
        fetchData()
    }, []);

    return (
        <div>
            <CntWrap title={'Fear&Greed Index'}>
                {
                    data.length ? data.map((d, i) => (
                        <div key={i}>
                            {d.now_value}
                        </div>
                    ))
                        : ('loading')
                }
            </CntWrap>
        </div>
    );
}

export default Fgi;
