import { useState } from 'react'

const Home = () => {

    const [count, setCount] = useState<number>(0)

    return (
        <>
            <div>This is the Home page</div>

            <div>
                <button onClick={() => setCount(count + 1)}>Count: {count}</button>
            </div>
        </>
    )


}

export default Home;