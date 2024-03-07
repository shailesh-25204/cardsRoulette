import React from 'react'
import { useSelector } from 'react-redux'

function Leaderboard() {
  const leaderboard = useSelector(state => state.leaderboard)
  console.log("this is leaderboard ", leaderboard)
  return (
    <div>{leaderboard.players.map((user,index) =>(
      <div key={index}>
        {`${user.username} ${user.rank}`}
      </div>
    ))}</div>
  )
}

export default Leaderboard