import Place from "./Place";

export default function Location() {
  return (
    <div className="location">
      <p className="heading">Choose your location</p>
      <div className="flex flex-wrap mx-auto justify-center items-end gap-8">
        <Place cityName="City" />
        <Place cityName="City" />
        <Place cityName="City" />
        <Place cityName="Cityyyyyyyyyyyyyyyyyyyyyyyy" />
        <Place cityName="City" />
      </div>
    </div>
  );
}
