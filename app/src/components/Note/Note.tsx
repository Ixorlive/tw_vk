// src/components/Note.js
import React from 'react';
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

function Note({ note, editable, onEdit }: { note: NoteType, editable: Boolean, onEdit: any }) {
    return (
        <div>
            <h4>User: {note.user_id}</h4>
            <p>{note.content}</p>
            <p>Created {note.created_at}</p>
            {editable && <button onClick={onEdit}>Edit</button>}
            {editable && <button onClick={() => { handleDeleteNote(note.id) }}>Delete</button>}
        </div>
    )
}


export default Note;