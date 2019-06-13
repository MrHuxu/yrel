import React, { Component } from 'react';
import { shape, arrayOf, string, number, any, object } from 'prop-types';
import { connect } from 'react-redux';

import Lexer from './Lexer';
import Parser from './Parser';
import Output from './Output';
import styles from '../styles/results';

class Results extends Component {
  static propTypes = {
    content : shape({
      outputs : arrayOf(string),
      tokens  : arrayOf(shape({
        lineNum  : number.isRequired,
        category : number.isRequired,
        value    : any.isRequired
      })),
      statements : arrayOf(object)
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

var mapStateToProps = function ({ result }) {
  return {
    content : result.content
  };
};

export default connect(mapStateToProps)(Results);
