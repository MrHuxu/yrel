import React, { Component } from 'react';
import { connect } from 'react-redux';
import Radium, { Style } from 'radium';

import { submitCode } from '../actions/EditorActions';
import styles from '../styles/editor';

@Radium
class Editor extends Component {
  componentDidMount () {
    this.editor = CodeMirror(this.refs.editorElem, {   // eslint-disable-line
      lineNumbers : true,
      value       :
`print "hello world\\n" * 3;
a = !true; b = false;

// this is a comment
if (3 > 1) {
	print a;
} else {
	print b;
}

b = 3 + 1;
print b;
print a / 0;

c = 4;
c = c - 1;
while (c > 0) {
	print c;
	c = c - 1;
}
print c;`,
      mode      : 'javascript',
      tabSize   : 2,
      autofocus : true
    });
  }

  _submit (e) {
    this.props.dispatch(submitCode(this.editor.getValue()));
  }

  render () {
    return (
      <div>
        <Style rules = {styles.global} />
        <h3 className = 'ui header'>Editor</h3>
        <div
          ref = 'editorElem'
          style = {styles.editorElem}
        />
        <div
          className = 'ui blue button'
          style = {styles.submitBtn}
          onClick = {this._submit.bind(this)}
        >
          Submit
        </div>
      </div>
    );
  }
}

export default connect()(Editor);
