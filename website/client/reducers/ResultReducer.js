import { REFRESH_RESULT } from '../actions/EditorActions';

export function result (state = {
  result  : '',
  content : {
    tokens     : [],
    statements : [],
    outputs    : []
  }
}, action) {
  const { type, content } = action;

  if (type === REFRESH_RESULT) {
    return Object.assign({}, content);
  }

  return state;
}
