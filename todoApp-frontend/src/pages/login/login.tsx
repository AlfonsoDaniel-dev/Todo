import { useState } from 'react'

const Login = () => {


    const [count, setCount] = useState<number>(0)

    const handleClick = () => {
        setCount(count + 1)
    }

    return (
        <>
            <h1>Auth</h1>
            <div>This is the Auth page</div>
            <button onClick={handleClick}>{count}</button>
        </>


    )
}

export default Login