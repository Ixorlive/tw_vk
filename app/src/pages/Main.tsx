import { useEffect, useState } from "react";
import { UserType } from "../entity/UserType";
import { NoteType } from "../entity/NoteType";
import { Header } from "../components/Header/Header";
import './Main.css';
import NoteList from "../components/NoteList/NoteList";
import { NOTE_API_GETALL } from "../config";

interface MainPageProps {
    user: UserType
}

const Main: React.FC<MainPageProps> = ({ user }) => {
    const [notes, setNotes] = useState<NoteType[]>([]);
    const [filter, setFilter] = useState<string>('all')
    const [timeFilter, setTimeFilter] = useState<number | null>(null)
    const currTime = new Date();

    useEffect(() => {
        const fetchAllNotes = async () => {
            try {
                const response = await fetch(NOTE_API_GETALL);
                if (!response.ok) {
                    throw new Error('Failed to fetch notes');
                }
                const data = await response.json() || [];

                setNotes(data);
            } catch (error) {
                console.error('Error fetching notes:', error);
            }
        };

        fetchAllNotes();
    }, []);

    const handleFilterChange = (newFilter: string) => {
        setFilter(newFilter);
    };

    const handleTimeFilterChange = (newTimeFilter: number | null) => {
        setTimeFilter(newTimeFilter);
    };

    const filteredNotes = notes.filter(note => {
        var cond = filter === 'all' || note.user_id === user.id;
        if (timeFilter !== null) {
            const createdAt = new Date(note.created_at);
            const dayDiff = (currTime.getTime() - createdAt.getTime()) / (1000 * 3600 * 24);
            cond &&= dayDiff <= timeFilter;
        }
        return cond
    })

    return (
        <div>
            <Header user={user} applyTabFilter={handleFilterChange} applyTimeFilter={handleTimeFilterChange} />
            <main>
                <NoteList user={user} notes={filteredNotes} showControlButtons={user.id !== 0 && filter === 'my'} />
            </main>
        </div>
    )
}

export default Main;