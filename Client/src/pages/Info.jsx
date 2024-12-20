import React from "react";

export default function Info() {
  return (
    <div className="info-page">
      <div className="flex">
        <div className="grid grid-cols-2 gap-0">
          <img src="/image 16.png" alt="" />
          <img src="/image 13.png" alt="" />
          <img src="/image 14.png" alt="" />
          <img src="/image 15.png" alt="" />
        </div>
        <div className="px-16 py-10 flex justify-center flex-col gap-3 w-1/2">
          <p className="heading">The future is flexible</p>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. In
            obcaecati ipsum delectus quidem magnam! Fugiat error, atque dolores
            harum unde consequatur facilis nobis a accusantium labore, assumenda
            iusto vero hic!
          </p>
        </div>
      </div>
    </div>
  );
}
