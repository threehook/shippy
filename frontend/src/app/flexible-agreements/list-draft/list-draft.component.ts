import { Component, OnInit } from '@angular/core';
// import { DraftItem } from '../models/agreement';
// import { DraftService } from '../services/draft.service';
import { MatDialog } from '@angular/material';
// import { NewDraftComponent } from '../new-draft/new-draft.component';

@Component({
  selector: 'app-list-draft',
  templateUrl: './list-draft.component.html',
  styleUrls: ['./list-draft.component.css']
})
export class ListDraftComponent implements OnInit {
  // drafts: DraftItem[] = [];

  // constructor(private draftService: DraftService, private dialog: MatDialog) { }

  ngOnInit() {
    this.loadDrafts();
  }

  private loadDrafts() {
    // this.draftService.getDrafts().subscribe((drafts) => { this.drafts = drafts; });
  }

  // openAddDialog() {
  //   const dialogRef = this.dialog.open(NewDraftComponent, {
  //     width: '400px'
  //   });
  //   dialogRef.afterClosed().subscribe(() => {
  //     this.loadDrafts();//GOTO draft
  //   })
  // }
  //
  // hasNoDrafts() {
  //   return this.drafts.length < 1;
  // }
}
