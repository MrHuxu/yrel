import React, { Component } from 'react';
import { arrayOf, shape, number, any } from 'prop-types';

import styles from '../styles/lexer';

class Lexer extends Component {
  static propTypes = {
    data : arrayOf(shape({
      lineNum  : number.isRequired,
      category : number.isRequired,
      value    : any.isRequired
    }))
  };

  render () {
    const { data } = this.props;
    const colors = [null, 'brown', 'orange', 'red', 'olive', 'green', 'teal', 'blue', 'purple', 'pink', 'grey', 'yellow', 'violet', 'black'];

    var lines = {};
    data.forEach((token) => {
      if (!lines[token.lineNum]) {
        lines[token.lineNum] = [token];
      } else {
        lines[token.lineNum].push(token);
      }
    });

    return (
      <div className = 'ui stacked segment'>
        <h4 className = 'ui horizontal divider header'>
          Lexer
        </h4>

        { data.length ? Object.keys(lines).map(line => {
          return (
            <div key = {`line-${line}`}>
              <a
                style = {styles.labelContainer}
                className = 'ui basic circular small label'
              >
                {`#${line}`}
              </a>
              {lines[line].map((token, index) => {
                return (
                  <a
                    key = {`line-${line}-token-${index}`}
                    style = {styles.labelContainer}
                    className = {`ui ${colors[token.category]} label`}
                  >
                    {`${token.value}`}
                  </a>
                );
              })}
            </div>
          );
        }) : <p>List all tokens here</p> }
      </div>
    );
  }
}

export default Lexer;
