import React from 'react';

import styles from '../styles/menu-bar';

export default function renderMenuBar (params) {
  return (
    <div
      className = 'ui inverted menu'
      style = {styles.bar}
    >
      <div
        className = 'left menu'
        style = {styles.leftBar}
      >
        <a className = 'item'>
          <h2> Yrel </h2>
        </a>
        <a className = 'item'>
          <h5> a script language written in Go </h5>
        </a>
      </div>

      <div
        className = 'right menu'
        style = {styles.rightBar}
      >
        <a className = 'active item'>
          Editor
        </a>
        <a className = 'item'>
          Spec
        </a>
        <a className = 'item'>
          About
        </a>
      </div>
    </div>
  );
};
