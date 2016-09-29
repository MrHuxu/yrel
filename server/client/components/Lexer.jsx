import React, { Component } from 'react';

class Lexer extends Component {
  render () {
    const { data } = this.props;
    return (
      <div className = 'ui stacked segment'>
        <h4 className = 'ui header'>Lexer</h4>
        { data.length ? data.map(output => <div>{output[0]}</div>) : <p>List all tokens here</p> }
      </div>
    );
  }
}

export default Lexer;
