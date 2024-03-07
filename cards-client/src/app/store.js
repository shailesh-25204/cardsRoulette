import {configureStore} from '@reduxjs/toolkit'
import wsReducer from '../features/websocketSlice'

export const store =configureStore ({
    reducer: wsReducer,
})