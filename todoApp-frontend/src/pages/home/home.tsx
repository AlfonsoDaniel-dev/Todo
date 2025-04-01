import { useState } from 'react'

import styles from "./styles/home.module.css"

const Home = () => {

    const [count, setCount] = useState<number>(0)

    return (
        <div className={styles.home}>

            <div>This is the Home page</div>

            <div>
                <button onClick={() => setCount(count + 1)}>Count: {count}</button>
            </div>

        </div>
    )


}

export default Home;