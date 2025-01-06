import { useParams } from "react-router";

export default function PropertyDetails() {
  const { propertyId } = useParams();
  console.log(propertyId);
  return <div>{propertyId}</div>;
}
