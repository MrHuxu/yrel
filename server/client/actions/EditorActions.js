export const REFRESH_RESULT = 'REFRESH_RESULT';
export function refreshResult (result) {
  return {
    type    : REFRESH_RESULT,
    content : result
  };
};

export function submitCode (code) {
  return function (dispatch) {
    var request = new Request('/yrel/', {
      method : 'POST',
      body   : code
    });
    fetch(request).then(res => {
      return res.json();
    }).then(json => {
      dispatch(refreshResult(json));
    });
  };
}
