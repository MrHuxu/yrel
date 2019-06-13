import React from 'react';
import { StyleRoot } from 'radium';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { rootStore } from './store';

import App from './components/App';

render(
  <StyleRoot>
    <Provider store = {rootStore}>
      <App />
    </Provider>
  </StyleRoot>,
  document.getElementById('yrel')
);
