import React, { useState, useEffect } from 'react';
import { NOTE_API_CREATE, NOTE_API_EDIT } from '../../config';
import { NoteType } from '../../entity/NoteType';
import { UserType } from '../../entity/UserType';

interface NoteFormProps {
    user?: UserType
    note?: NoteType

    onSave: any
}

const NoteForm: React.FC<NoteFormProps> = ({ user, note, onSave }) => {
    const [noteContent, setNoteContent] = useState('');
    const isNewNote = note === undefined;

    useEffect(() => {
        if (!isNewNote) {
            setNoteContent(note.content)
        }
    }, [note, isNewNote]);

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault();
        const url = isNewNote ? NOTE_API_CREATE : `${NOTE_API_EDIT}/${note.id}`;
        const method = isNewNote ? 'POST' : 'PUT';
        const user_id = isNewNote ? user!.id : note.user_id

        try {
            const response = await fetch(url, {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ content: noteContent, user_id: user_id }),
            });
            if (!response.ok) {
                throw new Error(isNewNote ? 'Failed to create note' : 'Failed to update note');
            }
            alert(isNewNote ? 'Note created successfully' : 'Note updated successfully');
            onSave()
        } catch (error) {
            console.error(isNewNote ? 'Error creating note:' : 'Error updating note:', error);
            alert(isNewNote ? 'Failed to create note' : 'Failed to update note');
            onSave()
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Note Content:
                <textarea value={noteContent} onChange={e => setNoteContent(e.target.value)} />
            </label>
            <button type="submit">{isNewNote ? 'Create Note' : 'Update Note'}</button>
        </form>
    );
}

export default NoteForm;
