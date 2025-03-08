
import { useParams } from "react-router-dom";

const Session = () => {
  const { id } = useParams();
  return (
    <div className="space-y-8 animate-fade-in">
      <h1 className="text-3xl font-bold tracking-tight">Session Details</h1>
    </div>
  );
};

export default Session;
