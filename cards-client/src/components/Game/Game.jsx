import React from 'react'
import Card from './Card'
import { useSelector } from 'react-redux'
import Leaderboard from '../Leaderboard/Leaderboard'

function Game() {
    const card = useSelector(state => state.gameState)
    // console.log("this is card ",card)
  return (
    <div className='flex flex-col-reverse justify-end h-full w-full md:flex-row md:h-full'>
        <div className='h-1/2 md:w-96'>
            <Leaderboard/>
        </div>
        <div className='h-72 md:h-auto w-full flex flex-row justify-center items-center'>
            <div className='w-full grid grid-cols-2 grid-flow-row'>
                <Card/>
                <div className='flex justify-center'>
                    <div className="text-white text-2xl h-36 w-24 absolute bg-slate-900 border-[2px] rounded-md text-center flex items-center justify-center">
                        {card.nextCard}
                    </div>
                </div>
            </div>
            <section className='md:w-96 hidden md:block'>
                <ul>
                    <li>Section</li>
                    <li>Section</li>
                    <li>Section</li>
                    <li>Section</li>
                    <li>Section</li>
                </ul>
            </section>
        </div>
    </div>
  )
}

export default Game