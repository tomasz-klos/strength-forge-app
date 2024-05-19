import DateCarousel from "@organisms/dateCarousel";

const Home: React.FC = () => {
  return (
    <section className="flex-1 flex flex-col gap-12">
      <h1 className="text-xl font-medium text-center">StrengthForge App</h1>
      <DateCarousel />
    </section>
  );
};

export default Home;
