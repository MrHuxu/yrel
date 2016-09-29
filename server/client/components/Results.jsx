import React, { Component } from 'react';

import Lexer from './Lexer';
import Parser from './Parser';
import Output from './Output';
import styles from '../styles/results';

class Results extends Component {
  render () {
    return (
      <div>
        <div style = {styles.cardContainer}>
          <Lexer />
        </div>
        <div style = {styles.cardContainer}>
          <Parser />
        </div>
        <div style = {styles.cardContainer}>
          <Output />
        </div>
      </div>
    );
  }
}

export default Results;
