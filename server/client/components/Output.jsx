import React, { Component } from 'react';

import styles from '../styles/output';

class Output extends Component {
  render () {
    const { data } = this.props;
    return (
      <div
        style = {styles.container}
        className = 'ui piled segment'
      >
        <h4 className = 'ui horizontal divider header'>
          <i className = 'print icon' />
          Output
        </h4>
        { data.length ? data.map(output => <div style = {styles.text}>{output}</div>) : <p>List all outputs here</p> }
      </div>
    );
  }
}

export default Output;
