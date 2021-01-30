import * as Actions from './actions';
import initialState from '../store/initialState';

export const UiStateReducer = (state = initialState.uiState, action) => {
    switch (action.type) {
        case Actions.MODAL_OPEN:
            return {
                ...action.payload,
            };
        case Actions.MODAL_CLOSE:
            return {
                ...action.payload,
            };
        default:
            return state;
    }
};
