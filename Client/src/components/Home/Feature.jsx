export default function Feature(props) {
  return (
    <div className="min-w-[260px] md:w-1/6 flex flex-col gap-2 bg-[#E2F1E8] p-8 rounded-3xl">
      <p className="-mx-1">{props.icon}</p>
      <p className="font-semibold text-2xl">{props.feature}</p>
      <p>{props.info}</p>
    </div>
  );
}
