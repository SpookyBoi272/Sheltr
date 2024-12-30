import SearchBar from "./SearchBar";

export default function Hero() {
  return (
    <div className="hero">
      <div
        className="main"
        style={{
          background: `url(/hero.png)`,
          backgroundRepeat: "no-repeat",
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        <div className="info">
          <p className="heading">We rent your property</p>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Error
            dolores, debitis iste cumque praesentium tempora labore perspiciatis
            perferendis soluta illum!
          </p>
        </div>
      </div>
      <div className="relative w-[70%] mx-auto">
        <SearchBar />
      </div>
    </div>
  );
}
