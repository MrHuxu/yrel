import React, { Component } from 'react';
import Radium, { Style } from 'radium';

import renderMenuBar from './MenuBar';
import Editor from './Editor';
import Results from './Results';
import styles from '../styles/app';

@Radium
class App extends Component {
  render () {
    return (
      <div>
        <Style rules = {styles.global} />
        <div>
          {renderMenuBar()}
        </div>
        <div style = {styles.leftPanel}>
          <Editor />
        </div>
        <div style = {styles.rightPanel}>
          <Results />
        </div>
      </div>
    );
  }
}

export default App;
