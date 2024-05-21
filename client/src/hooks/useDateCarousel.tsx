import { useState, useEffect, useCallback } from "react";
import { CarouselApi } from "@atoms/carousel";

interface CarouselState {
  api?: CarouselApi;
  currentIndex: number;
  dateRange: Date[];
}

interface CarouselHook {
  api?: CarouselApi;
  setApi: (api: CarouselApi) => void;
  currentIndex: number;
  dateRange: Date[];
  startIndex?: number;
  formatDate: (date: Date | null) => string;
}

const useDateCarousel = (): CarouselHook => {
  const [api, setApi] = useState<CarouselApi | undefined>(undefined);
  const [state, setState] = useState<CarouselState>({
    currentIndex: 0,
    dateRange: [],
  });

  const numDaysToShow = 11;
  const startIndex = Math.floor(numDaysToShow / 2);

  const { currentIndex, dateRange } = state;

  const generateDateRange = useCallback((numDaysToShow: number): Date[] => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - Math.floor(numDaysToShow / 2));

    const dates: Date[] = [];
    for (let i = 0; i < numDaysToShow; i++) {
      const currentDate = new Date(startDate);
      currentDate.setDate(startDate.getDate() + i);
      dates.push(currentDate);
    }
    return dates;
  }, []);

  useEffect(() => {
    const dates = generateDateRange(numDaysToShow);

    setState((prevState) => ({
      ...prevState,
      dateRange: dates,
      currentIndex: Math.floor(numDaysToShow / 2),
    }));
  }, [generateDateRange]);

  useEffect(() => {
    if (!api) return;

    const onSelect = (config: any) => {
      const index = config.selectedScrollSnap();
      setState((prevState) => ({ ...prevState, currentIndex: index }));
    };

    api.on("select", onSelect);

    return () => {
      api.off("select", onSelect);
    };
  }, [api]);

  const formatDate = (date: Date | null): string => {
    if (!date) return "";

    const today = new Date();
    const tomorrow = new Date(today);
    tomorrow.setDate(today.getDate() + 1);

    const yesterday = new Date(today);
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
