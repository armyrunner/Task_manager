
import DeskWork from '../assets/deskwork.jpg';      // your original image
import Nietjuhart from '../assets/nietjuhart.jpg';     // add 2 more real images
import RockStatue from '../assets/rockformation.jpg';

const carouselQuotes = [
  {
    quote: "Action is the foundational key to all success.",
    author: "Pablo Picasso"
  },
  {
    quote: "The secret of getting ahead is getting started. The secret of getting started is breaking your complex overwhelming tasks into small, manageable tasks, and then starting on the first one.",
    author: "Mark Twain"
  },
  {
    quote: "Productivity is never an accident. It is always the result of a commitment to excellence, intelligent planning, and focused effort.",
    author: "Paul J. Meyer"
  }
];

function Cover() {
  return (
    <div
      id="carouselExampleIndicators"
      className="carousel slide carousel-fade"
      data-bs-ride="carousel"
      data-bs-interval='15000'
    >
      {/* Indicators */}
      <div className="carousel-indicators">
        <button
          type="button"
          data-bs-target="#carouselExampleIndicators"
          data-bs-slide-to="0"
          className="active"
          aria-current="true"
          aria-label="Slide 1"
        />
        <button
          type="button"
          data-bs-target="#carouselExampleIndicators"
          data-bs-slide-to="1"
          aria-label="Slide 2"
        />
        <button
          type="button"
          data-bs-target="#carouselExampleIndicators"
          data-bs-slide-to="2"
          aria-label="Slide 3"
        />
      </div>

      {/* Slides with overlay text */}
      <div className="carousel-inner">
        {[DeskWork, Nietjuhart,RockStatue].map((imgSrc, index) => (
          <div
            key={index}
            className={`carousel-item ${index === 0 ? 'active' : ''}`}
          >
            <img
              src={imgSrc}
              className="d-block w-100 object-fit-cover"
              alt={`Workspace slide ${index + 1}`}
              style={{ maxHeight: '65vh',
                objectPosition: 'center',
              }}
            />
            {/* Quote overlay */}
            <div
              className="position-absolute top-50 start-50 translate-middle text-center text-white px-4 py-3"
              style={{
                backgroundColor: 'rgba(0, 0, 0, 0.45)', // semi-transparent dark bg for contrast
                borderRadius: '12px',
                maxWidth: '80%',
                textShadow: '0 2px 8px rgba(0,0,0,0.8)'
              }}
            >
              <h3 className="mb-2 fw-bold" style={{ fontSize: '2rem', color: '#ffffff' }}>
                "{carouselQuotes[index].quote}"
              </h3>
              <p className="mb-0 fst-italic" style={{ color: '#e0e0e0' }}>
                — {carouselQuotes[index].author}
              </p>
            </div>
          </div>
        ))}
      </div>

      {/* Controls */}
      <button
        className="carousel-control-prev"
        type="button"
        data-bs-target="#carouselExampleIndicators"
        data-bs-slide="prev"
      >
        <span className="carousel-control-prev-icon" aria-hidden="true" />
        <span className="visually-hidden">Previous</span>
      </button>
      <button
        className="carousel-control-next"
        type="button"
        data-bs-target="#carouselExampleIndicators"
        data-bs-slide="next"
      >
        <span className="carousel-control-next-icon" aria-hidden="true" />
        <span className="visually-hidden">Next</span>
      </button>
    </div>
  );
}

export default Cover;