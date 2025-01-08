import { useParams } from "react-router";

export default function PropertyDetails() {
  const { propertyId } = useParams();
  const amenities = [
    {name: "WiFi", icon: "/television.svg"},
    {name: "Hot Shower", icon: "/television.svg", id: 1},
    {name: "Hot Shower", icon: "/television.svg", id: 2},
    {name: "Hot Shower", icon: "/television.svg", id: 3},
    {name: "Hot Shower", icon: "/television.svg", id: 4},
  ]

  return <div className="min-w-96">

    <h1 className="heading">Rhoncus Suspendisse</h1>
    <p>London, Notthing</p>
    <h2>Description</h2>
    <p>
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Architecto aut placeat vel. Adipisci ad, dignissimos officia modi pariatur esse ratione perferendis labore ex distinctio non quo nulla dolorem. Similique, eum!
      Eaque maiores architecto, rerum fugiat optio est error odit ratione. Et eaque hic necessitatibus qui, temporibus explicabo ab eveniet voluptate animi similique voluptates corrupti, unde eos enim saepe perferendis maiores?
      Vel voluptatum iusto et? Unde libero molestias assumenda voluptas molestiae tenetur illo commodi enim recusandae tempora dolor ducimus labore delectus, nostrum dignissimos in sint id? Facilis rem illum error fuga?
    </p>
    <h2>Amenities</h2>
    <div className="amenities flex align-middle justify-center">
      {amenities.map((amenity) => {
        return <div key={amenity.id} className="amenity flex align-middle justify-center mr-4">
          <span>{amenity.name}</span>
          <span className="w-10 h-10">
            <img className="w-full h-full" src={amenity.icon} alt={amenity.name+" icon"} />
          </span>
        </div>
      })}
    </div>

    <div className="location">
      <h2>Location</h2>
      <div className="map w-9/12 h-64 border-2">

      </div>
    </div>

  </div>;
}
