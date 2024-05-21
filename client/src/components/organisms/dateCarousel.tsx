import * as React from "react";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "@atoms/carousel";
import useDateCarousel from "@hooks/useDateCarousel";

const DateCarousel: React.FC = () => {
  const { setApi, startIndex, currentIndex, dateRange, formatDate } =
    useDateCarousel();

  const renderCalendarSlides = () => {
    return dateRange.map((date, index) => (
      <CarouselItem key={index}>
        <div>{date.toLocaleDateString()}</div>
      </CarouselItem>
    ));
  };

  return (
    <Carousel
      className="flex-1  flex flex-col gap-4"
      orientation="horizontal"
      setApi={setApi}
      opts={{
        startIndex,
      }}
    >
      <div className="flex items-center justify-between">
        <CarouselPrevious className="static transform-none" />
        <p>{formatDate(dateRange[currentIndex])}</p>
        <CarouselNext className="static transform-none" />
      </div>

      <CarouselContent>{renderCalendarSlides()}</CarouselContent>
    </Carousel>
  );
};

export default DateCarousel;
