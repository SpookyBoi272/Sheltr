export default function SearchBar() {
  return (
    <div
      className="search-bar"
      // style={{ boxShadow: "inset 0 0 0 3px #064749" }}
    >
      <div className="">
        <button>Select a city</button>
      </div>
      <div className="">
        <button>Move in -- Move out</button>
      </div>
      <div className="">
        <button>Guests</button>
      </div>
      <div className="">
        <button className="bg-[#064749] font-semibold text-white rounded-[50px] px-8">
          Search
        </button>
      </div>
    </div>
  );
}
