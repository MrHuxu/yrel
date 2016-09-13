import React, { Component } from 'react';

import Tokens from './Tokens';
import Expressions from './Expressions';
import Output from './Output';
import styles from '../styles/results';

class Results extends Component {
  render () {
    return (
      <div>
        <div style = {styles.cardContainer}>
          <Tokens />
        </div>
        <div style = {styles.cardContainer}>
          <Expressions />
        </div>
        <div style = {styles.cardContainer}>
          <Output />
        </div>
      </div>
    );
  }
}

export default Results;
