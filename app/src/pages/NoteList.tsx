import React, { useEffect, useState } from 'react';
import { UserType } from '../entity/UserType'
import { NoteType } from '../entity/NoteType'
import Note from '../components/Note/Note';
import { NOTE_API_GETALL } from '../config';
import Modal from '../components/Modal/Modal';
import NoteForm from '../components/NoteForm/NoteForm';

interface NoteListProps {
    user: UserType
}

const NoteList: React.FC<NoteListProps> = ({ user }) => {
    const [activeTab, setActiveTab] = useState('All');
    const [notes, setNotes] = useState<NoteType[]>([]);
    const [loading, setLoading] = useState(false);
    const [showModal, setShowModal] = useState(false);
    const [currentNote, setCurrentNote] = useState<NoteType | null>(null);

    useEffect(() => {
        const fetchAllNotes = async () => {
            setLoading(true);
            try {
                const response = await fetch(NOTE_API_GETALL);
                if (!response.ok) {
                    throw new Error('Failed to fetch notes');
                }
                const data = await response.json();

                setNotes(data);
            } catch (error) {
                console.error('Error fetching notes:', error);
            } finally {
                setLoading(false);
            }
        };

        fetchAllNotes();
    }, []);

    const filteredNotes = activeTab === 'My' && user.id !== 0
        ? notes.filter(note => note.user_id === user.id)
        : notes;

    const handleEditClick = (note: NoteType) => {
        setCurrentNote(note);
        setShowModal(true);
    };

    const handleNewNoteClick = () => {
        setCurrentNote(null)
        setShowModal(true);
    };

    const closeModal = () => {
        setShowModal(false);
        setCurrentNote(null);
    };

    return (
        <div>
            <h1>List of notes</h1>
            <div>
                <button onClick={() => setActiveTab('All')}>All</button>
                {user.id !== 0 && (
                    <>
                        <button onClick={() => setActiveTab('My')}>My</button>
                        <button onClick={handleNewNoteClick}>+ New Note</button>
                    </>
                )}
            </div>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <div>
                    {filteredNotes.map(note => (
                        <Note note={note} editable={activeTab === 'My'} onEdit={() => handleEditClick(note)} />
                    ))}
                </div>
            )}
            {showModal && (
                <Modal onClose={closeModal}>
                    <NoteForm user={user} note={currentNote} onSave={closeModal} />
                </Modal>
            )}
        </div>
    );
};

export default NoteList;
