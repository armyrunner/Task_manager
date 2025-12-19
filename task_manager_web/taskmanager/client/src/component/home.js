

const Home = () => {

    const handleClick = () => {
        console.log('hello,taskers')
    }

    return (
        <div className="home">
            <h2>Home Page!!!</h2>
            <button onClick={handleClick}>Click Here</button>
        </div>
      );
}
 
export default Home;