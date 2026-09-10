import { configureStore } from '@reduxjs/toolkit'
import authReducer from './auth/authSlice'
import checkoutReducer from './checkout/checkoutSlice'
//configuracion necesaria para redux
// export default configureStore({
//   reducer: {
//     auth:authReducer,
//     checkout:checkoutReducer,
//   }
// })

// Configuración de la store
const store = configureStore({
  reducer: {
    auth: authReducer,
    checkout: checkoutReducer,
  }
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;