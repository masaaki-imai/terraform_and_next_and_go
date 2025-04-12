import { Dialog, DialogTitle, DialogContent, List, ListItem, ListItemText, IconButton, Box } from '@mui/material';
import CloseIcon from '@mui/icons-material/Close';

type ValidationModalProps = {
  open: boolean;
  onClose: () => void;
  errors: string[];
};

export default function ValidationModal({ open, onClose, errors }: ValidationModalProps) {
  return (
    <Dialog open={open} onClose={onClose} maxWidth="sm" fullWidth>
      <DialogTitle>
        入力エラー
        <IconButton
          aria-label="close"
          onClick={onClose}
          sx={{ position: 'absolute', right: 8, top: 8 }}
        >
          <CloseIcon />
        </IconButton>
      </DialogTitle>
      <DialogContent>
        <List>
          {errors.map((error, index) => (
            <ListItem key={index}>
              <ListItemText
                primary={
                  <Box sx={{
                    color: 'error.main',
                    mb: 1.5,
                    fontWeight: 'bold',
                    fontSize: '1rem',
                    display: 'flex',
                    alignItems: 'flex-start'
                  }}>
                    <span style={{ marginRight: '8px' }}>・</span>
                    {error}
                  </Box>
                }
              />
            </ListItem>
          ))}
        </List>
      </DialogContent>
    </Dialog>
  );
}