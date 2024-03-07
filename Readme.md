# Cards-Roulette

Cards-Roulette is a full-stack web application and a game where players test their luck and strategy with a deck of cards. The game is built with a frontend in React + Vite using Tailwind CSS with Redux for state management, while the backend is powered by Golang. Redis serves as the database, providing efficient storage and retrieval of game data.



## Technologies Used

- Frontend: React + Vite, Tailwind CSS , Redux
- Backend: Golang + Gin
- Database: Redis

## How to Play

To play Cards-Roulette, simply visit [cardsRoulette site](https://cardsroulette-1.onrender.com/), enter your username to begin. Use your luck to reveal all the cards without encountering a bomb.

## Gameplay

In Cards-Roulette, there are four types of cards:

- Skip: Skips the next card in the deck.
- Reset: Resets the game by shuffling the deck.
- Bomb: Ends the game if revealed without a defuse card.
- Defuse: Counters a bomb card and can accumulate for one game.

Each game starts with the player drawing five random cards from the deck. The objective is to reveal all the cards without encountering a bomb. If a bomb card is revealed, the game ends unless the player has a defuse card to counter it. Defuse cards can pile up throughout the game, providing a strategic advantage.

## Development

If you wish to contribute to Cards-Roulette or explore its codebase, follow these steps:

1. Clone the repository:
```bash
   git clone https://github.com/shailesh-25204/cardsRoulette.git
```

2. Install Dependencies for frontend and backend:
```bash
    cd cards-roulette/cards-client
    npm install
    cd ../cards-server
    go mod tidy
```

3.Configure Environment variables:

- `REDIS_ENDPOINT`: The host of the Redis database.
- `REDIS_PORT`: The port of the Redis database.
- `REDIS_PASS`: The password of the Redis database.

These environment variables are used to configure the server. They should be set according to the deployment environment to ensure proper functionality.


4. Start the development servers:

```bash
# Client
cd ../cards-client
npm run dev

# Server
cd ../cards-server
go run .

```

Open your browser and navigate to http://localhost:5173 to access the Cards-Roulette application.

## Contributions

Feel free to fork this repository and submit pull requests with your enhancements. We appreciate any improvements or bug fixes you can provide. Thank you for your support!
