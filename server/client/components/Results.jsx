import React, { Component, PropTypes } from 'react';
import { connect } from 'react-redux';

import Lexer from './Lexer';
import Parser from './Parser';
import Output from './Output';
import styles from '../styles/results';

class Results extends Component {
  static propTypes = {
    content : PropTypes.shape({
      outputs : PropTypes.arrayOf(PropTypes.string),
      tokens  : PropTypes.arrayOf(PropTypes.shape({
        lineNum  : PropTypes.number.isRequired,
        category : PropTypes.number.isRequired,
        value    : PropTypes.any.isRequired
      })),
      statements : PropTypes.arrayOf(PropTypes.object)
    })
  };

  render () {
    const { content } = this.props;
    return (
      <div>
        <div style = {styles.cardContainer}>
          <Output data = {content.outputs} />
        </div>
        <div style = {styles.cardContainer}>
          <Lexer data = {content.tokens} />
        </div>
        <div style = {styles.cardContainer}>
          <Parser data = {content.statements} />
        </div>
      </div>
    );
  }
}

var mapStateToProps = function (state) {
  return {
    content : state.result.content
  };
};

export default connect(mapStateToProps)(Results);
