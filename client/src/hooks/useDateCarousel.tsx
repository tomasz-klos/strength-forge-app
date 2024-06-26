import { useState, useEffect, useCallback } from "react";
import { CarouselApi } from "@atoms/carousel";

interface CarouselState {
  api?: CarouselApi;
  startIndex?: number;
  currentIndex: number;
  dateRange: Date[];
}

interface Config {
  selectedScrollSnap: () => number;
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
    startIndex: 0,
    currentIndex: 0,
    dateRange: [],
  });

  const { startIndex, currentIndex, dateRange } = state;

  const generateDateRange = useCallback((numDays: number): Date[] => {
    const startDate = new Date();
    startDate.setDate(startDate.getDate() - Math.floor(numDays / 2));

    return Array.from({ length: numDays }, (_, i) => {
      const date = new Date(startDate);
      date.setDate(startDate.getDate() + i);
      return date;
    });
  }, []);

  useEffect(() => {
    const numDaysToShow = 11;
    const dates = generateDateRange(numDaysToShow);

    setState((prevState) => ({
      ...prevState,
      dateRange: dates,
      startIndex: Math.floor(numDaysToShow / 2),
      currentIndex: Math.floor(numDaysToShow / 2),
    }));
  }, [generateDateRange]);

  useEffect(() => {
    if (!api) return;

    const handleSelect = (config: Config) => {
      const index = config.selectedScrollSnap();
      setState((prevState) => ({ ...prevState, currentIndex: index }));
    };

    api.on("select", handleSelect);

    return () => {
      api.off("select", handleSelect);
    };
  }, [api]);

  const today = new Date();
  const tomorrow = new Date(today);
  tomorrow.setDate(today.getDate() + 1);

  const yesterday = new Date(today);
  yesterday.setDate(today.getDate() - 1);

  const formatDate = (date: Date | null): string => {
    if (!date) return "";

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
