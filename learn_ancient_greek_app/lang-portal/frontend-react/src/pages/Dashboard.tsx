
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const Dashboard = () => {
  return (
    <div className="space-y-8 animate-fade-in">
      <div className="space-y-2">
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
        <p className="text-muted-foreground">
          Welcome to your language learning journey
        </p>
      </div>

      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <Card>
          <CardHeader>
            <CardTitle>Last Session</CardTitle>
          </CardHeader>
          <CardContent>
            <p>No recent sessions</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Progress Overview</CardTitle>
          </CardHeader>
          <CardContent>
            <p>Start learning to see your progress</p>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Quick Actions</CardTitle>
          </CardHeader>
          <CardContent>
            <div className="space-y-2">
              <button className="w-full bg-primary text-primary-foreground px-4 py-2 rounded-md hover:opacity-90 transition-opacity">
                Start New Session
              </button>
              <button className="w-full bg-secondary text-secondary-foreground px-4 py-2 rounded-md hover:opacity-90 transition-opacity">
                View Words
              </button>
            </div>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default Dashboard;
