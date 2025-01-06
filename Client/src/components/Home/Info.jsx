import React from "react";
import {
  MdCalendarMonth,
  MdMessage,
  MdOutlineWifi,
  MdChair,
} from "react-icons/md";
import Feature from "./Feature";

export default function Info() {
  return (
    <div className="info-page">
      <div className="flex flex-col lg:flex-row mb-24 rounded-[40px] lg:rounded-[50px] overflow-hidden bg-[#F2F0F2] ">
        <div className="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-2 gap-0 [&>*]:object-cover [&>*]:h-[100%] [&>*]:w-[100%]">
          <img src="/image 16.png" alt="people" />
          <img src="/image 13.png" alt="people" />
          <img src="/image 14.png" alt="people" />
          <img src="/image 15.png" alt="people" />
        </div>
        <div className="px-10 lg:px-16 py-8 lg:py-10 flex justify-center flex-col gap-3 lg:w-3/5">
          <p className="heading">The future is flexible</p>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. In
            obcaecati ipsum delectus quidem magnam! Fugiat error, atque dolores
            harum unde consequatur facilis nobis a accusantium labore, assumenda
            iusto vero hic!
          </p>
        </div>
      </div>
      <div>
        <div className="max-w-[80%] mx-auto text-center space-y-4">
          <p className="heading">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Quo,
            consectetur.
          </p>
          <p>
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Cupiditate,
            odio aperiam! A autem assumenda ullam ipsa delectus optio officiis
            suscipit at voluptatem! Natus doloribus voluptates dolore
            voluptatibus architecto, nesciunt hic.
          </p>
        </div>
        <div className="mt-12 flex flex-row flex-wrap gap-6 justify-center">
          <Feature
            feature="Felxible living"
            info="Stay as Long or as little as you need with month-to-
                month contracts"
            icon={<MdCalendarMonth size={50} />}
          />
          <Feature
            feature="Move-in ready"
            info="Stay as Long or as little as you need with month-to-
                month contracts"
            icon={<MdChair size={50} />}
          />
          <Feature
            feature="High-speed Wi-Fi"
            info="Stay as Long or as little as you need with month-to-
                month contracts"
            icon={<MdOutlineWifi size={50} />}
          />
          <Feature
            feature="24/7 support"
            info="Stay as Long or as little as you need with month-to-
                month contracts"
            icon={<MdMessage size={50} />}
          />
        </div>
      </div>
    </div>
  );
}
