import React from 'react'
import {useSelector} from 'react-redux'


function Header() {
  const connectionStatus =useSelector(state => state.connStatus)
  const username = useSelector(state => state.username)
  return (
    <div className='flex flex-col justify-center items-center'>
    <div className='text-5xl '>
        Header
    </div>
    <div className='text-2xl '>
    Welcome {username}, Status: {connectionStatus}
    </div>
    </div>
  )
}

export default Header