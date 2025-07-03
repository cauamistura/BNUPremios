import React from 'react';
import Slider from 'react-slick';
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

interface ImageCarouselProps {
  images: string[];
  altPrefix?: string;
}

const ImageCarousel: React.FC<ImageCarouselProps> = ({ images, altPrefix = 'carousel-image' }) => {
  return (
    <Slider dots={true} infinite={true} speed={500} slidesToShow={1} slidesToScroll={1} arrows={true}>
      {images.map((img, idx) => (
        <div key={idx} className="reward-carousel-image">
          <img src={img} alt={`${altPrefix} ${idx + 1}`} />
        </div>
      ))}
    </Slider>
  );
};

export default ImageCarousel; 