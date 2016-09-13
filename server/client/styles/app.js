export default {
  global : {
    '*' : {
      fontFamily    : "'Rambla', sans-serif",
      letterSpacing : '0.03em'
    },

    '.CodeMirror *, code' : {
      fontFamily : '"Monaco", "MonacoRegular", "Courier New", monospace !important'
    }
  },

  leftPanel : {
    position : 'fixed',
    display  : 'inline-block',
    padding  : '10% 0 0 2%',
    height   : '100%',
    width    : '45%',
    fontSize : '20px',
    top      : '0',
    left     : '0'
  },

  rightPanel : {
    display       : 'inline-block',
    // borderLeft    : '1px dotted gray',
    verticalAlign : 'top',
    margin        : '56px 0 0 47%',
    height        : '90%',
    width         : '53%'
  }
};
