import { useState, useEffect, useCallback } from "react";
import { CarouselApi } from "@atoms/carousel";

const useDateCarousel = () => {
  const [api, setApi] = useState<CarouselApi | undefined>();
  const [startIndex, setStartIndex] = useState<number>(0);
  const [currentIndex, setCurrentIndex] = useState<number>(0);
  const [dateRange, setDateRange] = useState<Date[]>([]);

  const generateDateRange = useCallback((numDaysToShow: number) => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - Math.floor(numDaysToShow / 2));

    const dates = [];
    for (let i = 0; i < numDaysToShow; i++) {
      const currentDate = new Date(startDate);
      currentDate.setDate(startDate.getDate() + i);
      dates.push(currentDate);
    }
    return dates;
  }, []);

  useEffect(() => {
    const numDaysToShow = 11;
    const dates = generateDateRange(numDaysToShow);

    setDateRange(dates);
    setStartIndex(Math.floor(numDaysToShow / 2));
    setCurrentIndex(Math.floor(numDaysToShow / 2));
  }, [generateDateRange]);

  useEffect(() => {
    if (!api) return;

    const onSelect = (config: any) => {
      const index = config.selectedScrollSnap();
      setCurrentIndex(index);
    };

    api.on("select", onSelect);

    return () => {
      api.off("select", onSelect);
    };
  }, [api]);

  const formatDate = (date: Date) => {
    if (!date) return "";

    const today = new Date();
    const tomorrow = new Date();
    const yesterday = new Date();

    tomorrow.setDate(today.getDate() + 1);
    yesterday.setDate(today.getDate() - 1);

    const isToday = date.toDateString() === today.toDateString();
    const isTomorrow = date.toDateString() === tomorrow.toDateString();
    const isYesterday = date.toDateString() === yesterday.toDateString();

    if (isToday) return "Today";
    if (isTomorrow) return "Tomorrow";
    if (isYesterday) return "Yesterday";

    return date.toLocaleDateString("en-US", {
      weekday: "long",
      month: "long",
      day: "numeric",
    });
  };

  return {
    api,
    setApi,
    startIndex,
    currentIndex,
    dateRange,
    formatDate,
  };
};

export default useDateCarousel;
