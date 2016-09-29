import React, { Component } from 'react';

class Output extends Component {
  render () {
    const { data } = this.props;
    return (
      <div className = 'ui stacked segment'>
        <h4 className = 'ui header'>Output</h4>
        { data.length ? data.map(output => <div>{output}</div>) : <p>List all outputs here</p> }
      </div>
    );
  }
}

export default Output;
