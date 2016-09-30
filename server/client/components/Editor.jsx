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
`greet = "hello yrel. ";

// print statement, this line only contains a comment
print greet;

// print greet 3 times, and this is another comment
print greet * 3;

// if-else statement
if (3 > 1) {
  print "3 is larger than 1.";
} else {
  print "3 is smaller than 1.";
}

// use while statement to get the 10th fibnacci number
a = 1;
b = 1;
count = 2;
while (count < 10) {
  tmp = b;
  b = a + b;
  a = tmp;
  count = count + 1;
}
print "the 10th fibonacci number is: " + b;`,
      mode      : 'javascript',
      tabSize   : 2,
      autofocus : true
    });
    this._submit();
  }

  _submit () {
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
