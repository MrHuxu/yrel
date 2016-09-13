import React from 'react';
import { StyleRoot } from 'radium';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { rootStore } from './store';
import injectTapEventPlugin from 'react-tap-event-plugin';

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

import App from './components/App';

injectTapEventPlugin();

render(
  <StyleRoot>
    <MuiThemeProvider muiTheme = {getMuiTheme()}>
      <Provider store = {rootStore}>
        <App />
      </Provider>
    </MuiThemeProvider>
  </StyleRoot>,
  document.getElementById('yrel')
);
