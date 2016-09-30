import React, { Component } from 'react';

import styles from '../styles/lexer';

class Lexer extends Component {
  render () {
    const { data } = this.props;
    const colors = [null, 'red', 'orange', 'yellow', 'olive', 'green', 'teal', 'blue', 'purple', 'pink', 'grey', 'brown', 'violet', 'black'];
    return (
      <div className = 'ui stacked segment'>
        <h4 className = 'ui header'>Lexer</h4>
        { data.length ? data.map(output => {
          return (
            <a
              style = {styles.labelContainer}
              className = {`ui ${colors[output.Category]} label`}
            >
              {output.Value}
            </a>
          );
        }) : <p>List all tokens here</p> }
      </div>
    );
  }
}

export default Lexer;
