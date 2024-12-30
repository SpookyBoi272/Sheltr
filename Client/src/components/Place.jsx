export default function Place(props) {
  return (
    <div className="w-1/6 min-w-[196px] flex flex-col items-center overflow-hidden">
      <p className="my-3 text-2xl font-semibold block mx-auto break-all">
        {props.cityName}
      </p>
      <div className="rounded-3xl overflow-hidden">
        <img src="/image 13.png" alt="" className="h-[248px] object-cover" />
      </div>
    </div>
  );
}
