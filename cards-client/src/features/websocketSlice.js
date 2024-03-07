import { createSlice } from "@reduxjs/toolkit";

const initialState ={
    //add game state,
    //Leaderboard state
    gameState: {
        type: "newGame",
        len: 6,
        nextCard: "Start Playing",
        result: false,
    },
    leaderboard: {
        type: "leaderboard",
        count: 0,
        players: [],
    },
    connStatus: "Closed",
    socketUrl: "wss://cardsroulette.onrender.com",
    wsJsonResponse: {
        type: "newGame",
        len: 6,
        nextCard: "Start Playing",
        result: false,
    },
    username: ""
    
}

export const wsSlice =createSlice({
    name: 'websocket',
    initialState,
    reducers: {
        updateStatus: (state,action) => {
            state.connStatus = action.payload
        },
        changeUrl: (state,action) => {
            state.socketUrl = action.payload
        },
        newJsonMsg: (state,action) => {
            state.wsJsonResponse = action.payload
        },
        updateUsername: (state,action) => {
            state.username = action.payload
        },
        updateGameState: (state,action) => {
            state.gameState = action.payload
        },
        updateLeaderboard: (state,action) => {
            state.leaderboard = action.payload
        }
    }
})

export const {updateLeaderboard,updateStatus,changeUrl,newJsonMsg, updateUsername, updateGameState} = wsSlice.actions

export default wsSlice.reducer