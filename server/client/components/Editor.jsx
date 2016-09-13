import React, { Component } from 'react';
import Radium from 'radium';

import styles from '../styles/editor';

@Radium
class Editor extends Component {
  componentDidMount () {
    CodeMirror(this.refs.editorElem, {   // eslint-disable-line
      lineNumbers : true,
      value       :
`sum = 0
i = 1
while i < 10 {
  sum = sum + i
  i = i + 1
}

if i % 2 == 0 {
  even = even + 1
} else {
  odd = odd + 1
}`,
      mode    : 'javascript',
      tabSize : 2
    });
  }

  render () {
    return (
      <div>
        <h3 className = 'ui header'>Editor</h3>
        <div
          ref = 'editorElem'
          style = {styles.editorElem}
        />
        <div
          className = 'ui blue button'
          style = {styles.submitBtn}
        >
          Submit
        </div>
      </div>
    );
  }
}

export default Editor;
