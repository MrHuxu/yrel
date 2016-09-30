import React, { Component } from 'react';

import styles from '../styles/lexer';

class Lexer extends Component {
  render () {
    const { data } = this.props;
    const colors = [null, 'brown', 'orange', 'red', 'olive', 'green', 'teal', 'blue', 'purple', 'pink', 'grey', 'yellow', 'violet', 'black'];

    var lines = {};
    data.forEach((token) => {
      if (!lines[token.LineNum])
        lines[token.LineNum] = [token];
      else
        lines[token.LineNum].push(token);
    });

    return (
      <div className = 'ui stacked segment'>
        <h4 className = 'ui horizontal divider header'>
          Lexer
        </h4>

        { data.length ? Object.keys(lines).map(line => {
          return (
            <div>
              <a
                style = {styles.labelContainer}
                className = 'ui basic circular small label'
              >
                {`#${line}`}
              </a>
              {lines[line].map(token => {
                return (
                  <a
                    style = {styles.labelContainer}
                    className = {`ui ${colors[token.Category]} label`}
                  >
                    {token.Value}
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
