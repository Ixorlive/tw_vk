// src/components/Note.js
import React from 'react';
import './Note.css'
import { NoteType } from '../../entity/NoteType';
import { NOTE_API_DELETE } from '../../config';

const handleDeleteNote = async (noteId: number) => {
    const url = `${NOTE_API_DELETE}/${noteId}`;
    try {
        const response = await fetch(url, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
            },
        });
        if (!response.ok) {
            const data = await response.json();
            throw new Error("Can't delete note: " + data.error);
        }
        alert("Note is removed");
    } catch (error) {
        console.error('Error deleting note', error);
        alert("Error deleting note =(");
    }
}

interface NoteProps {
    note: NoteType;
    editable: Boolean;
    onEdit: any;
}

const Note: React.FC<NoteProps> = ({ note, editable, onEdit }) => {
    return (
        <div className="note">
            <h2>User: {note.user_id}</h2>
            <p>{note.content}</p>
            <p>Created {note.created_at}</p>
            {editable && <button className='button small' onClick={onEdit}>Edit</button>}
            {editable && <button className='button small red' onClick={() => { handleDeleteNote(note.id) }}>Delete</button>}
        </div>
    )
}


export default Note;
