export default {
  global : {
    '.CodeMirror' : {
      height : '76%'
    },

    '.CodeMirror *, code' : {
      fontFamily : '"Monaco", "MonacoRegular", "Courier New", monospace !important'
    }
  },

  editorElem : {
    border       : '1px solid rgba(34,36,38,.15)',
    fontSize     : '.6em',
    borderRadius : '4px',

    ':focus' : {
      border    : '1px solid rgba(85,117,225,.15)',
      boxShadow : '.3px .3px .3px #32C1DD'
    }
  },

  submitBtn : {
    margin : '20px 0 0 0'
  }
};
