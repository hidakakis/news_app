import React from 'react';
import AppNavigator from './navigation/AppNavigator';
import { Provider } from 'react-redux';
import store, { persistor } from './store';
import { PersistGate } from 'redux-persist/integration/react';

const App = () => {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <AppNavigator />
      </PersistGate>
    </Provider>
  );
};

export default App;
