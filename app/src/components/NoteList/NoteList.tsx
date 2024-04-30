import { useState } from "react";
import { NoteType } from "../../entity/NoteType";
import Note from "../Note/Note";
import Modal from "../Modal/Modal";
import NoteForm from "../NoteForm/NoteForm";
import { UserType } from "../../entity/UserType";
import './NoteList.css'

interface NoteListProps {
    user: UserType
    notes: NoteType[]

    showControlButtons: boolean
}

const NoteList: React.FC<NoteListProps> = ({ user, notes, showControlButtons }) => {
    const [showModal, setShowModal] = useState(false);
    const [currentNote, setCurrentNote] = useState<NoteType | undefined>(undefined);

    const handleEditClick = (note: NoteType) => {
        setCurrentNote(note);
        setShowModal(true);
    };

    const closeModal = () => {
        setShowModal(false);
        setCurrentNote(undefined);
    };

    const handleNewNoteClick = () => {
        setCurrentNote(undefined)
        setShowModal(true);
    };
    return (
        <div>
            {user.id !== 0 && <button className="button add-note" onClick={handleNewNoteClick}>+ Note</button>}
            {notes.map(note => (
                <Note key={note.id} note={note} editable={showControlButtons} onEdit={() => handleEditClick(note)} />
            ))}
            {showModal && (
                <Modal onClose={closeModal}>
                    <NoteForm user={user} note={currentNote} onSave={closeModal} />
                </Modal>
            )}
        </div>
    )
}

export default NoteList